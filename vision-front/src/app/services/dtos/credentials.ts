export default class Credentials {
  constructor(
    private email: string,
    private password: string,
    private errors: string[] = [],
  ) { }

  public validate(): string[] {
    if (!this.email || this.email.trim() === '') {
      this.errors.push('Email é obrigatório');
    }

    if (this.email && !this.email.includes('@')) {
      this.errors.push('Email inválido');
    }

    if (this.email && this.email.length > 255) {
      this.errors.push('Email deve ter no máximo 255 caracteres');
    }

    if (!this.password || this.password.trim() === '') {
      this.errors.push('Senha é obrigatória');
    }

    if (this.password && (this.password.length < 6 || this.password.length > 255)) {
      this.errors.push('Senha deve ter entre 6 e 255 caracteres');
    }

    return this.errors;
  }

  public toJSON(): string {
    return JSON.stringify({
      email: this.email,
      password: this.password,
    });
  }
}
