export default class Project {
  account: string;

  title: string;

  description: string;

  imageLink: string;

  githubLink: string;

  demoLink: string;

  technologies: string[];

  constructor(
    account: string,
    title: string,
    description: string,
    githubLink: string,
    demoLink: string,
    imageLink: string,
    technologies: string[],
  ) {
    this.account = account;
    this.title = title;
    this.imageLink = imageLink;
    this.description = description;
    this.githubLink = githubLink;
    this.demoLink = demoLink;
    this.technologies = technologies;
  }

  validate() {
    const errors = [];

    if (!this.title || this.title.trim() === '') {
      errors.push('Título é obrigatório');
    }

    if (this.title && (this.title.length < 3 || this.title.length > 255)) {
      errors.push('Título deve ter entre 3 e 255 caracteres');
    }

    if (!this.description || this.description.trim() === '') {
      errors.push('Descrição é obrigatória');
    }

    if (this.description && (this.description.length < 3 || this.description.length > 2000)) {
      errors.push('Descrição deve ter entre 3 e 2000 caracteres');
    }

    if (!this.account || this.account.trim() === '') {
      errors.push('Conta é obrigatória');
    }

    if (this.githubLink && this.githubLink.length > 500) {
      errors.push('GitHub deve ter no máximo 500 caracteres');
    }

    if (this.demoLink && this.demoLink.length > 500) {
      errors.push('Demo deve ter no máximo 500 caracteres');
    }

    if (this.imageLink && this.imageLink.length > 500) {
      errors.push('Imagem deve ter no máximo 500 caracteres');
    }

    return errors;
  }

  public toJSON() {
    return JSON.stringify({
      accountId: this.account,
      title: this.title,
      description: this.description,
      githubLink: this.githubLink,
      demoLink: this.demoLink,
      imageLink: this.imageLink,
      technologies: this.technologies,
    });
  }
}
