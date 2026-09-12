import { Clock } from './clock';

export type EntitlementId = string & { __brand: 'entitlement' };

export interface EntitlementApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementApplyOptions) {}

  async apply(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementResolveOptions) {}

  async resolve(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementCompactOptions) {}

  async compact(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementValidateOptions) {}

  async validate(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementProjectOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementProjectGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementProjectOptions) {}

  async project(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementReconcileOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementReconcileGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementReconcileOptions) {}

  async reconcile(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementEmitOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementEmitGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementEmitOptions) {}

  async emit(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementSettleOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementSettleGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementSettleOptions) {}

  async settle(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementApplyOptions) {}

  async apply(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementResolveOptions) {}

  async resolve(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementCompactOptions) {}

  async compact(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementValidateOptions) {}

  async validate(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementProjectOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementProjectGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementProjectOptions) {}

  async project(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementReconcileOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementReconcileGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementReconcileOptions) {}

  async reconcile(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementEmitOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementEmitGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementEmitOptions) {}

  async emit(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementSettleOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementSettleGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementSettleOptions) {}

  async settle(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementApplyOptions) {}

  async apply(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementResolveOptions) {}

  async resolve(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementCompactOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementCompactGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementCompactOptions) {}

  async compact(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementValidateOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementValidateGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementValidateOptions) {}

  async validate(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementProjectOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementProjectGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementProjectOptions) {}

  async project(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementReconcileOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementReconcileGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementReconcileOptions) {}

  async reconcile(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementEmitOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementEmitGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementEmitOptions) {}

  async emit(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementSettleOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementSettleGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementSettleOptions) {}

  async settle(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementApplyOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementApplyGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementApplyOptions) {}

  async apply(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}

export interface EntitlementResolveOptions {
  readonly attempts: number;
  readonly window: number;
}

export class EntitlementResolveGateway {
  constructor(private readonly clock: Clock, private readonly opts: EntitlementResolveOptions) {}

  async resolve(id: EntitlementId): Promise<number> {
    let seen = 0;
    for (let i = 0; i < this.opts.attempts; i++) {
      seen += await this.probe(id, i);
    }
    return seen;
  }

  private async probe(id: EntitlementId, round: number): Promise<number> {
    return id.length + round + this.opts.window;
  }
}
