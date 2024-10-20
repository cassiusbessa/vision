package com.github.cassiusbessa.vision.domain.service.dtos.profile;

import java.util.List;
import java.util.UUID;

public class ProfileUpdateCommand {

    private final UUID profileId;
    private final String name;
    private final String title;
		private final String image;
    private final String description;
    private final List<UUID> technologies;
    private final UUID starProjectId;
    private final UUID accountId;
    private final String link;

    public ProfileUpdateCommand() {
        this.profileId = null;
        this.name = null;
        this.title = null;
				this.image = null;
        this.description = null;
        this.technologies = null;
        this.starProjectId = null;
        this.accountId = null;
        this.link = null;
    }

    public ProfileUpdateCommand(UUID profileId, String name, String title, String image, String description, List<UUID> technologies, UUID starProjectId, UUID accountId, String link) {
        this.profileId = profileId;
        this.name = name;
        this.title = title;
				this.image = image;
        this.description = description;
        this.technologies = technologies;
        this.starProjectId = starProjectId;
        this.accountId = accountId;
        this.link = link;
    }

    public UUID getProfileId() { return profileId; }

    public String getName() {
        return name;
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

    public List<UUID> getTechnologies() {
        return technologies;
    }

    public UUID getStarProjectId() {
        return starProjectId;
    }

    public UUID getAccountId() {
        return accountId;
    }

    public String getLink() {
        return link;
    }

}
