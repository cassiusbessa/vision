package com.github.cassiusbessa.vision.dataaccess.entities;

import jakarta.persistence.*;

import java.util.HashSet;
import java.util.Set;
import java.util.UUID;

@Entity
@Table(name = "profiles")
public class ProfileDataBaseEntity {

    @Id
    @Column(name = "id", nullable = false, unique = true)
    private UUID id;

    @Column(name = "name", nullable = false)
    private String name;

    @Column(name = "title", nullable = false)
    private String title;

    @Column(name = "description", nullable = false)
    private String description;

    @ManyToMany(cascade = CascadeType.ALL, fetch = FetchType.EAGER)
    @JoinTable(
            name = "profiles_tags",
            joinColumns = @JoinColumn(name = "profile_id"),
            inverseJoinColumns = @JoinColumn(name = "tag_id"))
    private Set<TagDataBaseEntity> technologies = new HashSet<>();

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "star_project", nullable = false)
    private ProjectDataBaseEntity starProject;


    @OneToOne
    @JoinColumn(name = "account_id", nullable = false)
    private AccountDataBaseEntity account;

    public ProfileDataBaseEntity() {
    }

    public ProfileDataBaseEntity(UUID id, String name, String title, String description, AccountDataBaseEntity account, ProjectDataBaseEntity starProject, Set<TagDataBaseEntity> technologies) {
        this.id = id;
        this.name = name;
        this.title = title;
        this.description = description;
        this.account = account;
        this.starProject = starProject;
        this.technologies = technologies;
    }

    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public Set<TagDataBaseEntity> getTechnologies() {
        return technologies;
    }

    public void setTechnologies(Set<TagDataBaseEntity> technologies) {
        this.technologies = technologies;
    }

    public ProjectDataBaseEntity getStarProject() {
        return starProject;
    }

    public void setStarProject(ProjectDataBaseEntity starProject) {
        this.starProject = starProject;
    }

    public AccountDataBaseEntity getAccount() {
        return account;
    }

    public void setAccount(AccountDataBaseEntity account) {
        this.account = account;
    }

}
