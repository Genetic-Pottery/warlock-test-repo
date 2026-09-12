import { Clock } from './clock';

export type DisputeId = string & { __brand: 'dispute' };

export interface DisputeApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeApplyOptions) {}

  async apply(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeResolveOptions) {}

  async resolve(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeCompactOptions) {}

  async compact(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeValidateOptions) {}

  async validate(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeProjectOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeProjectGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeProjectOptions) {}

  async project(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeReconcileOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeReconcileGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeReconcileOptions) {}

  async reconcile(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeEmitOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeEmitGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeEmitOptions) {}

  async emit(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeSettleOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeSettleGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeSettleOptions) {}

  async settle(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeApplyOptions) {}

  async apply(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeResolveOptions) {}

  async resolve(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeCompactOptions) {}

  async compact(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeValidateOptions) {}

  async validate(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeProjectOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeProjectGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeProjectOptions) {}

  async project(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeReconcileOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeReconcileGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeReconcileOptions) {}

  async reconcile(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeEmitOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeEmitGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeEmitOptions) {}

  async emit(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeSettleOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeSettleGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeSettleOptions) {}

  async settle(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeApplyOptions) {}

  async apply(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeResolveOptions) {}

  async resolve(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeCompactOptions) {}

  async compact(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeValidateOptions) {}

  async validate(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeProjectOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeProjectGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeProjectOptions) {}

  async project(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeReconcileOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeReconcileGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeReconcileOptions) {}

  async reconcile(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeEmitOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeEmitGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeEmitOptions) {}

  async emit(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeSettleOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeSettleGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeSettleOptions) {}

  async settle(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeApplyOptions) {}

  async apply(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface DisputeResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class DisputeResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: DisputeResolveOptions) {}

  async resolve(id: DisputeId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: DisputeId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}
