import { Clock } from './clock';

export type RollupId = string & { __brand: 'rollup' };

export interface RollupApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class RollupApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: RollupApplyOptions) {}

  async apply(id: RollupId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: RollupId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface RollupResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class RollupResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: RollupResolveOptions) {}

  async resolve(id: RollupId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: RollupId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface RollupCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class RollupCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: RollupCompactOptions) {}

  async compact(id: RollupId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: RollupId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface RollupValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class RollupValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: RollupValidateOptions) {}

  async validate(id: RollupId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: RollupId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface RollupProjectOptions {
  readonly attempts: number;
  readonly window: number;
}

export class RollupProjectGateway {
  constructor(private readonly clock: Clock, private readonly opts: RollupProjectOptions) {}

  async project(id: RollupId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: RollupId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface RollupReconcileOptions {
  readonly attempts: number;
  readonly window: number;
}

export class RollupReconcileGateway {
  constructor(private readonly clock: Clock, private readonly opts: RollupReconcileOptions) {}

  async reconcile(id: RollupId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: RollupId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface RollupEmitOptions {
  readonly attempts: number;
  readonly window: number;
}

export class RollupEmitGateway {
  constructor(private readonly clock: Clock, private readonly opts: RollupEmitOptions) {}

  async emit(id: RollupId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: RollupId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface RollupSettleOptions {
  readonly attempts: number;
  readonly window: number;
}

export class RollupSettleGateway {
  constructor(private readonly clock: Clock, private readonly opts: RollupSettleOptions) {}

  async settle(id: RollupId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: RollupId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface RollupApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class RollupApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: RollupApplyOptions) {}

  async apply(id: RollupId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: RollupId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface RollupResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class RollupResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: RollupResolveOptions) {}

  async resolve(id: RollupId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: RollupId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface RollupCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class RollupCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: RollupCompactOptions) {}

  async compact(id: RollupId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: RollupId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface RollupValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class RollupValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: RollupValidateOptions) {}

  async validate(id: RollupId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: RollupId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}
