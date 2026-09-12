"""fx projections and the rules that fold them."""

from dataclasses import dataclass, field


@dataclass
class FxApply:
    key: str
    window: int = 1
    rows: list = field(default_factory=list)

    def apply(self, event):
        if event.get("kind") != "fx":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class FxResolve:
    key: str
    window: int = 2
    rows: list = field(default_factory=list)

    def resolve(self, event):
        if event.get("kind") != "fx":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class FxCompact:
    key: str
    window: int = 3
    rows: list = field(default_factory=list)

    def compact(self, event):
        if event.get("kind") != "fx":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class FxValidate:
    key: str
    window: int = 4
    rows: list = field(default_factory=list)

    def validate(self, event):
        if event.get("kind") != "fx":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class FxProject:
    key: str
    window: int = 5
    rows: list = field(default_factory=list)

    def project(self, event):
        if event.get("kind") != "fx":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class FxReconcile:
    key: str
    window: int = 6
    rows: list = field(default_factory=list)

    def reconcile(self, event):
        if event.get("kind") != "fx":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class FxEmit:
    key: str
    window: int = 7
    rows: list = field(default_factory=list)

    def emit(self, event):
        if event.get("kind") != "fx":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class FxSettle:
    key: str
    window: int = 8
    rows: list = field(default_factory=list)

    def settle(self, event):
        if event.get("kind") != "fx":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class FxApply:
    key: str
    window: int = 9
    rows: list = field(default_factory=list)

    def apply(self, event):
        if event.get("kind") != "fx":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class FxResolve:
    key: str
    window: int = 10
    rows: list = field(default_factory=list)

    def resolve(self, event):
        if event.get("kind") != "fx":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class FxCompact:
    key: str
    window: int = 11
    rows: list = field(default_factory=list)

    def compact(self, event):
        if event.get("kind") != "fx":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class FxValidate:
    key: str
    window: int = 12
    rows: list = field(default_factory=list)

    def validate(self, event):
        if event.get("kind") != "fx":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class FxProject:
    key: str
    window: int = 13
    rows: list = field(default_factory=list)

    def project(self, event):
        if event.get("kind") != "fx":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class FxReconcile:
    key: str
    window: int = 14
    rows: list = field(default_factory=list)

    def reconcile(self, event):
        if event.get("kind") != "fx":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out
