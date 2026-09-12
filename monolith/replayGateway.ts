import { Clock } from './clock';

export type ReplayId = string & { __brand: 'replay' };

export interface ReplayApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ReplayApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: ReplayApplyOptions) {}

  async apply(id: ReplayId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ReplayId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ReplayResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ReplayResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: ReplayResolveOptions) {}

  async resolve(id: ReplayId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ReplayId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ReplayCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ReplayCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: ReplayCompactOptions) {}

  async compact(id: ReplayId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ReplayId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ReplayValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ReplayValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: ReplayValidateOptions) {}

  async validate(id: ReplayId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ReplayId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ReplayProjectOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ReplayProjectGateway {
  constructor(private readonly clock: Clock, private readonly opts: ReplayProjectOptions) {}

  async project(id: ReplayId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ReplayId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ReplayReconcileOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ReplayReconcileGateway {
  constructor(private readonly clock: Clock, private readonly opts: ReplayReconcileOptions) {}

  async reconcile(id: ReplayId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ReplayId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ReplayEmitOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ReplayEmitGateway {
  constructor(private readonly clock: Clock, private readonly opts: ReplayEmitOptions) {}

  async emit(id: ReplayId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ReplayId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ReplaySettleOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ReplaySettleGateway {
  constructor(private readonly clock: Clock, private readonly opts: ReplaySettleOptions) {}

  async settle(id: ReplayId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ReplayId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ReplayApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ReplayApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: ReplayApplyOptions) {}

  async apply(id: ReplayId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ReplayId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ReplayResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ReplayResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: ReplayResolveOptions) {}

  async resolve(id: ReplayId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ReplayId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ReplayCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ReplayCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: ReplayCompactOptions) {}

  async compact(id: ReplayId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ReplayId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ReplayValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ReplayValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: ReplayValidateOptions) {}

  async validate(id: ReplayId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ReplayId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}
