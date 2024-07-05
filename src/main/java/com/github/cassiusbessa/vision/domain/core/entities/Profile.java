package com.github.cassiusbessa.vision.domain.core.entities;

import com.github.cassiusbessa.vision.domain.core.valueobjects.AccountId;
import com.github.cassiusbessa.vision.domain.core.valueobjects.ProfileId;
import com.github.cassiusbessa.vision.domain.core.valueobjects.ProjectId;
import com.github.cassiusbessa.vision.domain.core.valueobjects.TagId;

import java.util.ArrayList;
import java.util.List;

public class Profile extends AggregateRoot<ProfileId> {

    private final String name;
    private final String title;
    private final String description;
    private final List<TagId> technologies;
    private final ProjectId starProject;
    private final AccountId accountId;
    private final List<String> failureMessages = new ArrayList<>();

    public Profile(ProfileId id, String name, String title, String description, List<TagId> technologies, ProjectId starProject, AccountId accountId) {
        super.setId(id);
        this.name = name;
        this.title = title;
        this.description = description;
        this.technologies = technologies;
        this.starProject = starProject;
        this.accountId = accountId;

    }

    public static Builder builder() {
        return new Builder();
    }

    public String getName() {
        return name;
    }

    public String getTitle() {
        return title;
    }

    public String getDescription() {
        return description;
    }

    public List<TagId> getTechnologies() {
        return technologies;
    }

    public ProjectId getStarProject() {
        return starProject;
    }

    public AccountId getAccountId() {
        return accountId;
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

        if (accountId == null){
            failureMessages.add("Account must be informed");
        }

    }

    public static final class Builder {
        private ProfileId id;
        private String name;
        private String title;
        private String description;
        private List<TagId> technologies;
        private ProjectId starProject;
        private AccountId accountId;

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

        public Builder withTitle(String title) {
            this.title = title;
            return this;
        }

        public Builder withDescription(String description) {
            this.description = description;
            return this;
        }

        public Builder withTechnologies(List<TagId> technologies) {
            this.technologies = technologies;
            return this;
        }

        public Builder withStarProject(ProjectId starProject) {
            this.starProject = starProject;
            return this;
        }

        public Builder withAccountId(AccountId accountId) {
            this.accountId = accountId;
            return this;
        }

        public Profile build() {
            return new Profile(id, name, title, description, technologies, starProject, accountId);
        }
    }

}
