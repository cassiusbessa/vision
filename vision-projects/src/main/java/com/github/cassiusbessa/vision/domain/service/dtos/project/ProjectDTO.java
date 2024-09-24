package com.github.cassiusbessa.vision.domain.service.dtos.project;

import com.github.cassiusbessa.vision.domain.core.valueobjects.ProjectLinks;
import com.github.cassiusbessa.vision.domain.service.dtos.tag.TagDTO;

import java.util.Date;
import java.util.Set;
import java.util.UUID;

public record ProjectDTO (UUID id, UUID accountId, String title, String image, String description, ProjectLinks links, Set<TagDTO> technologies, Date createdAt) { }
