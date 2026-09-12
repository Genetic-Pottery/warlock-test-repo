import { Clock } from './clock';

export type CouponId = string & { __brand: 'coupon' };

export interface CouponApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class CouponApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: CouponApplyOptions) {}

  async apply(id: CouponId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: CouponId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface CouponResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class CouponResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: CouponResolveOptions) {}

  async resolve(id: CouponId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: CouponId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface CouponCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class CouponCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: CouponCompactOptions) {}

  async compact(id: CouponId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: CouponId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface CouponValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class CouponValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: CouponValidateOptions) {}

  async validate(id: CouponId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: CouponId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface CouponProjectOptions {
  readonly attempts: number;
  readonly window: number;
}

export class CouponProjectGateway {
  constructor(private readonly clock: Clock, private readonly opts: CouponProjectOptions) {}

  async project(id: CouponId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: CouponId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface CouponReconcileOptions {
  readonly attempts: number;
  readonly window: number;
}

export class CouponReconcileGateway {
  constructor(private readonly clock: Clock, private readonly opts: CouponReconcileOptions) {}

  async reconcile(id: CouponId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: CouponId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface CouponEmitOptions {
  readonly attempts: number;
  readonly window: number;
}

export class CouponEmitGateway {
  constructor(private readonly clock: Clock, private readonly opts: CouponEmitOptions) {}

  async emit(id: CouponId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: CouponId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface CouponSettleOptions {
  readonly attempts: number;
  readonly window: number;
}

export class CouponSettleGateway {
  constructor(private readonly clock: Clock, private readonly opts: CouponSettleOptions) {}

  async settle(id: CouponId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: CouponId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface CouponApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class CouponApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: CouponApplyOptions) {}

  async apply(id: CouponId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: CouponId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface CouponResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class CouponResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: CouponResolveOptions) {}

  async resolve(id: CouponId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: CouponId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface CouponCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class CouponCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: CouponCompactOptions) {}

  async compact(id: CouponId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: CouponId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface CouponValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class CouponValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: CouponValidateOptions) {}

  async validate(id: CouponId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: CouponId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}
