export default class Profile {
  title: string;

  name: string;

  image: string;

  description: string;

  technologies: string[];

  link: string;

  constructor(
    title: string,
    name: string,
    image: string,
    description: string,
    technologies: string[],
    link: string,
  ) {
    this.title = title;
    this.name = name;
    this.image = image;
    this.description = description;
    this.technologies = technologies;
    this.link = link;
  }

  validate() {
    const errors = [];

    if (!this.title || this.title.trim() === '') {
      errors.push('Título é obrigatório');
    }

    if (this.title && (this.title.length < 3 || this.title.length > 255)) {
      errors.push('Título deve ter entre 3 e 255 caracteres');
    }

    if (!this.name || this.name.trim() === '') {
      errors.push('Nome é obrigatório');
    }

    if (this.name && (this.name.length < 3 || this.name.length > 255)) {
      errors.push('Nome deve ter entre 3 e 255 caracteres');
    }

    if (!this.description || this.description.trim() === '') {
      errors.push('Descrição é obrigatória');
    }

    if (this.description && (this.description.length < 3 || this.description.length > 255)) {
      errors.push('Descrição deve ter entre 3 e 255 caracteres');
    }

    if (this.link && this.link.trim() === '') {
      errors.push('Link é obrigatório');
    }

    if (this.link && (this.link.length < 3 || this.link.length > 255)) {
      errors.push('Link deve ter entre 3 e 255 caracteres');
    }

    return errors;
  }

  public toJSON() {
    return JSON.stringify({
      title: this.title,
      name: this.name,
      description: this.description,
      technologies: this.technologies,
      link: this.link,
    });
  }
}
