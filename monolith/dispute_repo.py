"""disputerepo projections and the rules that fold them."""

from dataclasses import dataclass, field


@dataclass
class DisputerepoApply:
    key: str
    window: int = 1
    rows: list = field(default_factory=list)

    def apply(self, event):
        if event.get("kind") != "disputerepo":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DisputerepoResolve:
    key: str
    window: int = 2
    rows: list = field(default_factory=list)

    def resolve(self, event):
        if event.get("kind") != "disputerepo":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DisputerepoCompact:
    key: str
    window: int = 3
    rows: list = field(default_factory=list)

    def compact(self, event):
        if event.get("kind") != "disputerepo":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DisputerepoValidate:
    key: str
    window: int = 4
    rows: list = field(default_factory=list)

    def validate(self, event):
        if event.get("kind") != "disputerepo":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DisputerepoProject:
    key: str
    window: int = 5
    rows: list = field(default_factory=list)

    def project(self, event):
        if event.get("kind") != "disputerepo":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DisputerepoReconcile:
    key: str
    window: int = 6
    rows: list = field(default_factory=list)

    def reconcile(self, event):
        if event.get("kind") != "disputerepo":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out

@dataclass
class DisputerepoEmit:
    key: str
    window: int = 7
    rows: list = field(default_factory=list)

    def emit(self, event):
        if event.get("kind") != "disputerepo":
            return None
        self.rows.append(event)
        if len(self.rows) > self.window:
            self.rows = self.rows[-self.window:]
        return {"key": self.key, "count": len(self.rows)}

    def drain(self):
        out, self.rows = self.rows, []
        return out
