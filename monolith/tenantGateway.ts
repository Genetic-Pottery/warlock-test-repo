import { Clock } from './clock';

export type TenantId = string & { __brand: 'tenant' };

export interface TenantApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class TenantApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: TenantApplyOptions) {}

  async apply(id: TenantId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: TenantId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface TenantResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class TenantResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: TenantResolveOptions) {}

  async resolve(id: TenantId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: TenantId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface TenantCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class TenantCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: TenantCompactOptions) {}

  async compact(id: TenantId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: TenantId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface TenantValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class TenantValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: TenantValidateOptions) {}

  async validate(id: TenantId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: TenantId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface TenantProjectOptions {
  readonly attempts: number;
  readonly window: number;
}

export class TenantProjectGateway {
  constructor(private readonly clock: Clock, private readonly opts: TenantProjectOptions) {}

  async project(id: TenantId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: TenantId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface TenantReconcileOptions {
  readonly attempts: number;
  readonly window: number;
}

export class TenantReconcileGateway {
  constructor(private readonly clock: Clock, private readonly opts: TenantReconcileOptions) {}

  async reconcile(id: TenantId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: TenantId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface TenantEmitOptions {
  readonly attempts: number;
  readonly window: number;
}

export class TenantEmitGateway {
  constructor(private readonly clock: Clock, private readonly opts: TenantEmitOptions) {}

  async emit(id: TenantId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: TenantId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface TenantSettleOptions {
  readonly attempts: number;
  readonly window: number;
}

export class TenantSettleGateway {
  constructor(private readonly clock: Clock, private readonly opts: TenantSettleOptions) {}

  async settle(id: TenantId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: TenantId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface TenantApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class TenantApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: TenantApplyOptions) {}

  async apply(id: TenantId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: TenantId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface TenantResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class TenantResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: TenantResolveOptions) {}

  async resolve(id: TenantId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: TenantId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface TenantCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class TenantCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: TenantCompactOptions) {}

  async compact(id: TenantId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: TenantId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface TenantValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class TenantValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: TenantValidateOptions) {}

  async validate(id: TenantId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: TenantId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}
