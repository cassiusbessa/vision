package com.github.cassiusbessa.vision.dataaccess.entities;

import jakarta.persistence.*;

import java.util.HashSet;
import java.util.Set;
import java.util.UUID;

@Entity
@Table(name = "profiles", indexes = {
	@Index(name = "idx_profile_link", columnList = "link")
	})
public class ProfileDataBaseEntity {

    @Id
    @Column(name = "id", nullable = false, unique = true)
    private UUID id;

    @Column(name = "name", nullable = false)
    private String name;

    @Column(name = "image")
    private String image;

    @Column(name = "title", nullable = false)
    private String title;

    @Column(name = "description", nullable = false, columnDefinition="varchar(1000)")
    private String description;

    @ManyToMany(cascade = CascadeType.ALL, fetch = FetchType.EAGER)
    @JoinTable(
            name = "profiles_tags",
            joinColumns = @JoinColumn(name = "profile_id"),
            inverseJoinColumns = @JoinColumn(name = "tag_id"))
    private Set<TagDataBaseEntity> technologies = new HashSet<>();

    @ManyToOne(fetch = FetchType.EAGER)
    @JoinColumn(name = "star_project", nullable = true)
    private ProjectDataBaseEntity starProject;


    @OneToOne
    @JoinColumn(name = "account_id", nullable = false)
    private AccountDataBaseEntity account;

    @Column(name = "link", unique = true)
    private String link;

    public ProfileDataBaseEntity() {
    }

    public ProfileDataBaseEntity(UUID id, String name, String image, String title, String description, AccountDataBaseEntity account, ProjectDataBaseEntity starProject, Set<TagDataBaseEntity> technologies, String link) {
        this.id = id;
        this.name = name;
        this.image = image;
        this.title = title;
        this.description = description;
        this.account = account;
        this.starProject = starProject;
        this.technologies = technologies;
        this.link = link;
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

    public String getImage() {
        return image;
    }

    public void setImage(String image) {
        this.image = image;
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

    public String getLink() { return link; }

    public void setLink(String link) { this.link = link; }

}
