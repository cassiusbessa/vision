package com.github.cassiusbessa.vision.domain.core.entities;

import com.github.cassiusbessa.vision.domain.core.valueobjects.ProfileId;


import java.util.ArrayList;
import java.util.List;
import java.util.Set;

public class Profile extends AggregateRoot<ProfileId> {

    private final String name;
    private final String image;
    private final String title;
    private final String description;
    private final Set<Tag> technologies;
    private final Project starProject;
    private final Account account;
    private final List<String> failureMessages = new ArrayList<>();

    public Profile(ProfileId id, String name, String image, String title, String description, Set<Tag> technologies, Project starProject, Account account) {
        super.setId(id);
        this.name = name;
        this.image = image;
        this.title = title;
        this.description = description;
        this.technologies = technologies;
        this.starProject = starProject;
        this.account = account;

    }

    public static Builder builder() {
        return new Builder();
    }

    public String getName() {
        return name;
    }

    public String getImage() {
        return image;
    }

    public String getTitle() {
        return title;
    }

    public String getDescription() {
        return description;
    }

    public Set<Tag> getTechnologies() {
        return technologies;
    }

    public Project getStarProject() {
        return starProject;
    }

    public Account getAccount() {
        return account;
    }

    public List<String> getFailureMessages() {
        return failureMessages;
    }

    public String getFailureMessagesAsString() {
        return String.join(", ", failureMessages);
    }

    public void validate() {
        if (name.length() < 3 || name.length() > 50 || name.isBlank()){
            failureMessages.add("Name must be between 3 and 50 characters");
        }

        if (title.length() < 3 || title.length() > 50 || title.isBlank()){
            failureMessages.add("Title must be between 3 and 100 characters");
        }

        if (description.length() < 3 || description.length() > 500 || description.isBlank()){
            failureMessages.add("Description must be between 3 and 500 characters");
        }

        if (technologies.size() > 6){
            failureMessages.add("Technologies must be less than 6");
        }

        if (account == null){
            failureMessages.add("Account must be informed");
        }

    }

    public static final class Builder {
        private ProfileId id;
        private String name;
        private String image;
        private String title;
        private String description;
        private Set<Tag> technologies;
        private Project starProject;
        private Account account;

        private Builder() {
        }

        public Builder withId(ProfileId id) {
            this.id = id;
            return this;
        }

        public Builder withName(String name) {
            this.name = name;
            return this;
        }

        public Builder withImage(String image) {
            this.image = image;
            return this;
        }

        public Builder withTitle(String title) {
            this.title = title;
            return this;
        }

        public Builder withDescription(String description) {
            this.description = description;
            return this;
        }

        public Builder withTechnologies(Set<Tag> technologies) {
            this.technologies = technologies;
            return this;
        }

        public Builder withStarProject(Project starProject) {
            this.starProject = starProject;
            return this;
        }

        public Builder withAccount(Account account) {
            this.account = account;
            return this;
        }

        public Profile build() {
            return new Profile(id, name, image, title, description, technologies, starProject, account);
        }
    }

}
