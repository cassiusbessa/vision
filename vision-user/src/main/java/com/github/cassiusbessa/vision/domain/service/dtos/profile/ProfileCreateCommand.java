package com.github.cassiusbessa.vision.domain.service.dtos.profile;

import java.util.List;
import java.util.UUID;

public class ProfileCreateCommand {

    private final String name;
    private final String title;
    private final String description;
    private final List<UUID> technologies;
    private  UUID accountId;
    private final String link;

    public ProfileCreateCommand() {
        this.link = null;
        this.name = null;
        this.title = null;
        this.description = null;
        this.technologies = null;
        this.accountId = null;
    }

    public ProfileCreateCommand(String name, String title, String description, List<UUID> technologies, UUID accountId, String link) {
        this.name = name;
        this.title = title;
        this.description = description;
        this.technologies = technologies;
        this.accountId = accountId;
        this.link = link;
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

    public List<UUID> getTechnologies() {
        return technologies;
    }

    public UUID getAccountId() {
        return accountId;
    }

    public String getLink() { return link; }

		public void setAccountId(UUID accountId) {
				this.accountId = accountId;
		}

}