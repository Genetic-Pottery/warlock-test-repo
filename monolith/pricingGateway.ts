import { Clock } from './clock';

export type PricingId = string & { __brand: 'pricing' };

export interface PricingApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class PricingApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: PricingApplyOptions) {}

  async apply(id: PricingId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: PricingId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface PricingResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class PricingResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: PricingResolveOptions) {}

  async resolve(id: PricingId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: PricingId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface PricingCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class PricingCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: PricingCompactOptions) {}

  async compact(id: PricingId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: PricingId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface PricingValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class PricingValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: PricingValidateOptions) {}

  async validate(id: PricingId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: PricingId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface PricingProjectOptions {
  readonly attempts: number;
  readonly window: number;
}

export class PricingProjectGateway {
  constructor(private readonly clock: Clock, private readonly opts: PricingProjectOptions) {}

  async project(id: PricingId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: PricingId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface PricingReconcileOptions {
  readonly attempts: number;
  readonly window: number;
}

export class PricingReconcileGateway {
  constructor(private readonly clock: Clock, private readonly opts: PricingReconcileOptions) {}

  async reconcile(id: PricingId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: PricingId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface PricingEmitOptions {
  readonly attempts: number;
  readonly window: number;
}

export class PricingEmitGateway {
  constructor(private readonly clock: Clock, private readonly opts: PricingEmitOptions) {}

  async emit(id: PricingId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: PricingId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface PricingSettleOptions {
  readonly attempts: number;
  readonly window: number;
}

export class PricingSettleGateway {
  constructor(private readonly clock: Clock, private readonly opts: PricingSettleOptions) {}

  async settle(id: PricingId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: PricingId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface PricingApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class PricingApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: PricingApplyOptions) {}

  async apply(id: PricingId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: PricingId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface PricingResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class PricingResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: PricingResolveOptions) {}

  async resolve(id: PricingId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: PricingId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface PricingCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class PricingCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: PricingCompactOptions) {}

  async compact(id: PricingId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: PricingId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface PricingValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class PricingValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: PricingValidateOptions) {}

  async validate(id: PricingId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: PricingId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}
