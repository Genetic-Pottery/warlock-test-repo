import { Clock } from './clock';

export type ShardId = string & { __brand: 'shard' };

export interface ShardApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ShardApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: ShardApplyOptions) {}

  async apply(id: ShardId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ShardId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ShardResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ShardResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: ShardResolveOptions) {}

  async resolve(id: ShardId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ShardId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ShardCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ShardCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: ShardCompactOptions) {}

  async compact(id: ShardId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ShardId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ShardValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ShardValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: ShardValidateOptions) {}

  async validate(id: ShardId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ShardId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ShardProjectOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ShardProjectGateway {
  constructor(private readonly clock: Clock, private readonly opts: ShardProjectOptions) {}

  async project(id: ShardId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ShardId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ShardReconcileOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ShardReconcileGateway {
  constructor(private readonly clock: Clock, private readonly opts: ShardReconcileOptions) {}

  async reconcile(id: ShardId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ShardId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ShardEmitOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ShardEmitGateway {
  constructor(private readonly clock: Clock, private readonly opts: ShardEmitOptions) {}

  async emit(id: ShardId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ShardId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ShardSettleOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ShardSettleGateway {
  constructor(private readonly clock: Clock, private readonly opts: ShardSettleOptions) {}

  async settle(id: ShardId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ShardId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ShardApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ShardApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: ShardApplyOptions) {}

  async apply(id: ShardId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ShardId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ShardResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ShardResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: ShardResolveOptions) {}

  async resolve(id: ShardId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ShardId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ShardCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ShardCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: ShardCompactOptions) {}

  async compact(id: ShardId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ShardId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface ShardValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class ShardValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: ShardValidateOptions) {}

  async validate(id: ShardId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: ShardId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}
