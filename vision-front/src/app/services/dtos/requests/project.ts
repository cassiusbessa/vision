export default class Project {
  id: string;

  account: string;

  title: string;

  image: string;

  description: string;

  links: string[];

  technologies: string[];

  createdAt: Date;

  constructor(
    id: string,
    account: string,
    title: string,
    image: string,
    description: string,
    links: string[],
    technologies: string[],
    createdAt: Date,
  ) {
    this.id = id;
    this.account = account;
    this.title = title;
    this.image = image;
    this.description = description;
    this.links = links;
    this.technologies = technologies;
    this.createdAt = createdAt;
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

    if (this.links && this.links.length === 0) {
      errors.push('Pelo menos um link é obrigatório');
    }

    return errors;
  }

  public toJSON() {
    return JSON.stringify({
      id: this.id,
      account: this.account,
      title: this.title,
      image: this.image,
      description: this.description,
      links: this.links,
      technologies: this.technologies,
      createdAt: this.createdAt,
    });
  }
}
