#!/usr/bin/env python3
"""Checks the outputs of `run.sh <lang> up` against sdk-e2e/expectations.json and compares languages.

Usage: check.py            -> writes sdk-e2e/.out/results.json, prints a summary, exit 1 on any failure
       check.py --self-test
Failures (exit 1): a missing output, a failed check, a `same` output that differs between languages.
Divergences (reported, no exit 1): xenorchestra resource state that differs between languages,
minus the paths matched by expectations.json "stateIgnore".
"""

import json
import re
import sys
from pathlib import Path

E2E = Path(__file__).resolve().parent.parent
TYPES = {"string": str, "number": (int, float), "boolean": bool, "array": list, "object": dict, "null": type(None)}


def check_value(value, rule):
    """Returns the list of failed checks for one output value."""
    errors = []
    t = rule.get("type")
    # bool is an int in python: do not let True pass as a number
    if t and (not isinstance(value, TYPES[t]) or (t == "number" and isinstance(value, bool))):
        errors.append(f"type: expected {t}, got {type(value).__name__}")
    if "equals" in rule and value != rule["equals"]:
        errors.append(f"equals: expected {rule['equals']!r}, got {value!r}")
    if "regex" in rule and not (isinstance(value, str) and re.search(rule["regex"], value)):
        errors.append(f"regex: {rule['regex']!r} does not match {value!r}")
    if rule.get("nonEmpty") and value in (None, "", [], {}):
        errors.append("nonEmpty: value is empty")
    return errors


def flatten(obj, prefix=""):
    """{"a": {"b": [1]}} -> {"a.b[0]": 1}"""
    if isinstance(obj, dict):
        out = {}
        for k, v in obj.items():
            out.update(flatten(v, f"{prefix}.{k}" if prefix else k))
        return out
    if isinstance(obj, list):
        out = {}
        for i, v in enumerate(obj):
            out.update(flatten(v, f"{prefix}[{i}]"))
        return out or {prefix: []}
    return {prefix: obj}


def resource_state(state, ignore):
    """xenorchestra resources of a `pulumi stack export`, keyed by type/name, flattened, ignored paths removed."""
    res = {}
    for r in state.get("deployment", {}).get("resources", []):
        typ = r.get("type", "")
        if not typ.startswith("xenorchestra:") or typ.startswith("pulumi:providers:"):
            continue
        key = f"{typ}/{r['urn'].rsplit('::', 1)[-1]}"
        flat = flatten(r.get("outputs", {}))
        res[key] = {p: v for p, v in flat.items() if not p.startswith("__") and not any(re.search(i, p) for i in ignore)}
    return res


def compare(expect, outputs, states):
    """outputs/states: {lang: dict}. Returns {"failures": [...], "divergences": [...]}."""
    failures, divergences = [], []
    rules = expect["outputs"]
    for lang, out in sorted(outputs.items()):
        for name, rule in rules.items():
            if name not in out:
                failures.append({"lang": lang, "output": name, "error": "missing output"})
                continue
            for err in check_value(out[name], rule):
                failures.append({"lang": lang, "output": name, "error": err})
        for name in sorted(set(out) - set(rules)):
            divergences.append({"lang": lang, "output": name, "error": "exported but not in expectations.json"})
    for name, rule in rules.items():
        values = {lang: out[name] for lang, out in outputs.items() if name in out}
        if rule.get("same") and len({json.dumps(v, sort_keys=True) for v in values.values()}) > 1:
            failures.append({"lang": "*", "output": name, "error": f"differs between languages: {values}"})

    flat = {lang: resource_state(s, expect.get("stateIgnore", [])) for lang, s in states.items()}
    for res in sorted({k for f in flat.values() for k in f}):
        per_lang = {lang: f.get(res) for lang, f in flat.items()}
        for lang, v in per_lang.items():
            if v is None:
                divergences.append({"lang": lang, "resource": res, "error": "resource missing in this language"})
        present = {lang: v for lang, v in per_lang.items() if v is not None}
        for path in sorted({p for v in present.values() for p in v}):
            values = {lang: v.get(path, "<absent>") for lang, v in present.items()}
            if len({json.dumps(x, sort_keys=True) for x in values.values()}) > 1:
                divergences.append({"resource": res, "path": path, "values": values})
    return {"failures": failures, "divergences": divergences}


def load_runs():
    outputs, states, statuses = {}, {}, {}
    for d in sorted((E2E / ".out").glob("*/")):
        if (d / "status").exists():
            statuses[d.name] = (d / "status").read_text().split("\n")
        if (d / "outputs.json").exists():
            outputs[d.name] = json.loads((d / "outputs.json").read_text())
        if (d / "state.json").exists():
            states[d.name] = json.loads((d / "state.json").read_text())
    return outputs, states, statuses


def main():
    expect = json.loads((E2E / "expectations.json").read_text())
    outputs, states, statuses = load_runs()
    result = compare(expect, outputs, states)
    result["steps"] = {lang: [s for s in st if s] for lang, st in statuses.items()}
    for lang, steps in result["steps"].items():
        if any(s.endswith("FAIL") for s in steps):
            result["failures"].append({"lang": lang, "error": "step failed: " + ", ".join(steps)})
    (E2E / ".out" / "results.json").write_text(json.dumps(result, indent=2))

    print("languages run:", ", ".join(f"{lang} ({' / '.join(s)})" for lang, s in result["steps"].items()) or "none")
    for f in result["failures"]:
        print("FAIL", json.dumps(f))
    for d in result["divergences"]:
        print("DIVERGENCE", json.dumps(d))
    print(f"{len(result['failures'])} failure(s), {len(result['divergences'])} divergence(s) -> {E2E / '.out/results.json'}")
    return 1 if result["failures"] or not outputs else 0


def self_test():
    assert check_value("x", {"type": "string", "nonEmpty": True}) == []
    assert check_value(True, {"type": "number"})
    assert check_value(1.0, {"equals": 1}) == []
    assert check_value("", {"nonEmpty": True})
    assert check_value("vm-pulumi-e2e-go", {"regex": "pulumi-e2e-"}) == []
    assert flatten({"a": {"b": [1, {"c": 2}]}, "d": []}) == {"a.b[0]": 1, "a.b[1].c": 2, "d": []}
    expect = {"outputs": {"id": {"nonEmpty": True}, "cpus": {"equals": 1, "same": True}}, "stateIgnore": [r"(^|\.)id$"]}
    st = lambda cpus, vid: {"deployment": {"resources": [
        {"type": "pulumi:providers:xenorchestra", "urn": "u::default", "outputs": {"x": 1}},
        {"type": "xenorchestra:index/vm:Vm", "urn": "u::vm", "outputs": {"id": vid, "cpus": cpus}}]}}
    r = compare(expect, {"go": {"id": "a", "cpus": 1}, "py": {"id": "b", "cpus": 1}}, {"go": st(1, "a"), "py": st(1, "b")})
    assert r == {"failures": [], "divergences": []}, r
    r = compare(expect, {"go": {"id": "a", "cpus": 1}, "py": {"cpus": 2, "extra": 0}}, {"go": st(1, "a"), "py": st(2, "b")})
    errors = sorted(f["error"].split(":")[0] for f in r["failures"])
    assert errors == ["differs between languages", "equals", "missing output"], r
    assert [d.get("path") or d.get("output") for d in r["divergences"]] == ["extra", "cpus"], r
    print("self-test ok")


if __name__ == "__main__":
    if sys.argv[1:] == ["--self-test"]:
        self_test()
    else:
        sys.exit(main())
