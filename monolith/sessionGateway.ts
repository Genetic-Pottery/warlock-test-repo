import { Clock } from './clock';

export type SessionId = string & { __brand: 'session' };

export interface SessionApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SessionApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: SessionApplyOptions) {}

  async apply(id: SessionId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SessionId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SessionResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SessionResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: SessionResolveOptions) {}

  async resolve(id: SessionId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SessionId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SessionCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SessionCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: SessionCompactOptions) {}

  async compact(id: SessionId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SessionId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SessionValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SessionValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: SessionValidateOptions) {}

  async validate(id: SessionId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SessionId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SessionProjectOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SessionProjectGateway {
  constructor(private readonly clock: Clock, private readonly opts: SessionProjectOptions) {}

  async project(id: SessionId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SessionId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SessionReconcileOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SessionReconcileGateway {
  constructor(private readonly clock: Clock, private readonly opts: SessionReconcileOptions) {}

  async reconcile(id: SessionId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SessionId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SessionEmitOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SessionEmitGateway {
  constructor(private readonly clock: Clock, private readonly opts: SessionEmitOptions) {}

  async emit(id: SessionId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SessionId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SessionSettleOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SessionSettleGateway {
  constructor(private readonly clock: Clock, private readonly opts: SessionSettleOptions) {}

  async settle(id: SessionId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SessionId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SessionApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SessionApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: SessionApplyOptions) {}

  async apply(id: SessionId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SessionId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SessionResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SessionResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: SessionResolveOptions) {}

  async resolve(id: SessionId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SessionId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SessionCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SessionCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: SessionCompactOptions) {}

  async compact(id: SessionId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SessionId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface SessionValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class SessionValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: SessionValidateOptions) {}

  async validate(id: SessionId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: SessionId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}
