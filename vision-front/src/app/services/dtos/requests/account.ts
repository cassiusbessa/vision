export default class Account {
  constructor(
    private name: string,
    private email: string,
    private password: string,
    private erros: string[] = [],
  ) { }

  public validate(): string[] {
    if (!this.name || this.name.trim() === '') {
      this.erros.push('Name is required');
    }

    if (this.name && (this.name.length < 3 || this.name.length > 255)) {
      this.erros.push('Name must be between 3 and 255 characters');
    }

    if (!this.email || this.email.trim() === '') {
      this.erros.push('Email is required');
    }

    if (this.email && !this.email.includes('@')) {
      this.erros.push('Email is invalid');
    }

    if (this.email && this.email.length > 255) {
      this.erros.push('Email must be less than 255 characters');
    }

    if (!this.password || this.password.trim() === '') {
      this.erros.push('Password is required');
    }

    if (this.password && (this.password.length < 6 || this.password.length > 255)) {
      this.erros.push('Password must be between 6 and 255 characters');
    }

    return this.erros;
  }

  public toJSON(): string {
    return JSON.stringify({
      email: this.email,
      password: this.password,
    });
  }
}
