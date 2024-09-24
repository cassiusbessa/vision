package com.github.cassiusbessa.vision.dataaccess.entities;

import jakarta.persistence.*;
import org.hibernate.annotations.OnDelete;
import org.hibernate.annotations.OnDeleteAction;

import java.util.Date;
import java.util.HashSet;
import java.util.Set;
import java.util.UUID;

@Entity
@Table(name = "projects")
public class ProjectDataBaseEntity {

    @Id
    @Column(name = "id", nullable = false, unique = true)
    private UUID id;

    @Column(name = "title", nullable = false)
    private String title;

    @Column(name = "image")
    private String image;

    @Column(name = "description", nullable = false)
    private String description;

    @Column(name = "repository")
    private String repositoryLink;

    @Column(name = "demo")
    private String demoLink;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "account_id", nullable = false)
    private AccountDataBaseEntity account;

    @Column(name = "created_at", nullable = false)
    private Date createdAt;

    @ManyToMany(fetch = FetchType.EAGER)
    @OnDelete(action = OnDeleteAction.CASCADE)
    @JoinTable(
            name = "projects_tags",
            joinColumns = @JoinColumn(name = "project_id"),
            inverseJoinColumns = @JoinColumn(name = "tag_id"))
    private Set<TagDataBaseEntity> technologies = new HashSet<>();

    public ProjectDataBaseEntity() {}

    public ProjectDataBaseEntity(UUID id, String title, String image, String description, String repositoryLink, AccountDataBaseEntity account, Date createdAt, Set<TagDataBaseEntity> technologies) {
        this.id = id;
        this.title = title;
        this.image = image;
        this.description = description;
        this.repositoryLink = repositoryLink;
        this.account = account;
        this.createdAt = createdAt;
        this.technologies = technologies;
    }

    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getImage() {
        return image;
    }

    public void setImage(String image) {
        this.image = image;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getRepositoryLink() {
        return repositoryLink;
    }

    public void setRepositoryLink(String repository) {
        this.repositoryLink = repository;
    }

    public String getDemoLink() {
        return demoLink;
    }

    public void setDemoLink(String demo) {
        this.demoLink = demo;
    }

    public AccountDataBaseEntity getAccount() {
        return account;
    }

    public void setAccount(AccountDataBaseEntity account) {
        this.account = account;
    }

    public Date getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(Date createdAt) {
        this.createdAt = createdAt;
    }

    public Set<TagDataBaseEntity> getTechnologies() {
        return technologies;
    }

    public void setTechnologies(Set<TagDataBaseEntity> technologies) {
        this.technologies = technologies;
    }

}
