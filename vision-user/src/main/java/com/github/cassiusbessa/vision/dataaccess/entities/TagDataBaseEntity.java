package com.github.cassiusbessa.vision.dataaccess.entities;

import jakarta.persistence.*;

import java.util.HashSet;
import java.util.Set;
import java.util.UUID;

@Entity
@Table(name = "tags")
public class TagDataBaseEntity {

    @Id
    @Column(name = "id", nullable = false, unique = true)
    private UUID id;

    @Column(name = "name", nullable = false)
    private String name;

    @ManyToMany(mappedBy = "technologies")
    private Set<ProfileDataBaseEntity> profiles = new HashSet<>();

    @ManyToMany(mappedBy = "technologies")
    private Set<ProjectDataBaseEntity> projects = new HashSet<>();


    public TagDataBaseEntity() {
    }

    public TagDataBaseEntity(UUID id, String name) {
        this.id = id;
        this.name = name;
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

    public Set<ProfileDataBaseEntity> getProfiles() { return profiles; }

    public void setProfiles(Set<ProfileDataBaseEntity> profiles) { this.profiles = profiles; }

    public Set<ProjectDataBaseEntity> getProjects() { return projects; }

    public void setProjects(Set<ProjectDataBaseEntity> projects) { this.projects = projects; }
}
