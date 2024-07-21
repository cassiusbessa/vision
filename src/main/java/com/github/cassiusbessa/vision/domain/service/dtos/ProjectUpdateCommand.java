package com.github.cassiusbessa.vision.domain.service.dtos;

import java.util.List;
import java.util.UUID;

public record ProjectUpdateCommand(UUID projectId, UUID accountId, String title, String description, String githubLink, String demoLink, String imageLink,
                                   List<UUID> technologies) {}
