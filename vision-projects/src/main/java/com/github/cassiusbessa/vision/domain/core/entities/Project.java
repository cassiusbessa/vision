package com.github.cassiusbessa.vision.domain.core.entities;

import com.github.cassiusbessa.vision.domain.core.valueobjects.ProjectId;
import com.github.cassiusbessa.vision.domain.core.valueobjects.ProjectLinks;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Set;

public class Project extends BaseEntity<ProjectId>{

    private final Account account;

    private final String title;

    private final String image;

    private final String description;

    private final ProjectLinks links;

    private final Set<Tag> technologies;

    private final Date createdAt;

    private final List<String> failureMessages = new ArrayList<>();


    public Project(ProjectId id, Account account, String title, String image, String description, ProjectLinks links, Set<Tag> technologies, Date createdAt) {
        super.setId(id);
        this.account = account;
        this.title = title;
        this.image = image;
        this.description = description;
        this.links = links;
        this.technologies = technologies;
        this.createdAt = createdAt;
    }

    public Account getAccount() {
        return account;
    }

    public String getTitle() {
        return title;
    }

    public String getImage() {
        return image;
    }

    public String getDescription() {
        return description;
    }

    public ProjectLinks getLinks() {
        return links;
    }

    public Set<Tag> getTechnologies() {
        return technologies;
    }

    public Date getCreatedAt() {
        return createdAt;
    }

    public List<String> getFailureMessages() {
        return failureMessages;
    }

    public String getFailureMessagesAsString() {
        return String.join(", ", failureMessages);
    }

    public void validate() {
        if (account == null) {
            failureMessages.add("Account is required");
        }
        if (title == null || title.isEmpty()) {
            failureMessages.add("Title is required");
        }
        if (description == null || description.isEmpty()) {
            failureMessages.add("Description is required");
        }
        if (technologies.size() > 6){
            failureMessages.add("Technologies must be less than 6");
        }
        if (createdAt == null) {
            failureMessages.add("Creation date is required");
        }
    }

    public static Builder builder() {
        return new Builder();
    }

    public static final class Builder {
        private ProjectId id;
        private Account account;
        private String title;
        private String image;
        private String description;
        private ProjectLinks links;
        private Set<Tag> technologies;
        private Date createdAt;


        public Builder withId(ProjectId id) {
            this.id = id;
            return this;
        }

        public Builder withAccount(Account account) {
            this.account = account;
            return this;
        }

        public Builder withTitle(String title) {
            this.title = title;
            return this;
        }

        public Builder withImage(String image) {
            this.image = image;
            return this;
        }

        public Builder withDescription(String description) {
            this.description = description;
            return this;
        }

        public Builder withLinks(ProjectLinks links) {
            this.links = links;
            return this;
        }

        public Builder withTechnologies(Set<Tag> technologies) {
            this.technologies = technologies;
            return this;
        }

        public Builder withCreatedAt(Date createdAt) {
            this.createdAt = createdAt;
            return this;
        }

        public Project build() {
            return new Project(id, account, title, image, description, links, technologies, createdAt);
        }
    }
}
