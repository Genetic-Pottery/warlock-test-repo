"""dunning projections and the rules that fold them."""

from dataclasses import dataclass, field


@dataclass
class DunningApply:
    key: str
    window: int = 1
    rows: list = field(default_factory=list)

    def apply(self, event):
        if event.get("kind") != "dunning":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DunningResolve:
    key: str
    window: int = 2
    rows: list = field(default_factory=list)

    def resolve(self, event):
        if event.get("kind") != "dunning":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DunningCompact:
    key: str
    window: int = 3
    rows: list = field(default_factory=list)

    def compact(self, event):
        if event.get("kind") != "dunning":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DunningValidate:
    key: str
    window: int = 4
    rows: list = field(default_factory=list)

    def validate(self, event):
        if event.get("kind") != "dunning":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DunningProject:
    key: str
    window: int = 5
    rows: list = field(default_factory=list)

    def project(self, event):
        if event.get("kind") != "dunning":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DunningReconcile:
    key: str
    window: int = 6
    rows: list = field(default_factory=list)

    def reconcile(self, event):
        if event.get("kind") != "dunning":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DunningEmit:
    key: str
    window: int = 7
    rows: list = field(default_factory=list)

    def emit(self, event):
        if event.get("kind") != "dunning":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DunningSettle:
    key: str
    window: int = 8
    rows: list = field(default_factory=list)

    def settle(self, event):
        if event.get("kind") != "dunning":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DunningApply:
    key: str
    window: int = 9
    rows: list = field(default_factory=list)

    def apply(self, event):
        if event.get("kind") != "dunning":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DunningResolve:
    key: str
    window: int = 10
    rows: list = field(default_factory=list)

    def resolve(self, event):
        if event.get("kind") != "dunning":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DunningCompact:
    key: str
    window: int = 11
    rows: list = field(default_factory=list)

    def compact(self, event):
        if event.get("kind") != "dunning":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DunningValidate:
    key: str
    window: int = 12
    rows: list = field(default_factory=list)

    def validate(self, event):
        if event.get("kind") != "dunning":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DunningProject:
    key: str
    window: int = 13
    rows: list = field(default_factory=list)

    def project(self, event):
        if event.get("kind") != "dunning":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DunningReconcile:
    key: str
    window: int = 14
    rows: list = field(default_factory=list)

    def reconcile(self, event):
        if event.get("kind") != "dunning":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out
