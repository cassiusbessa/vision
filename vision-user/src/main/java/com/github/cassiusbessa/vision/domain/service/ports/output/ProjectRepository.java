package com.github.cassiusbessa.vision.domain.service.ports.output;

import com.github.cassiusbessa.vision.domain.core.entities.Project;

import java.util.UUID;

public interface ProjectRepository {

    Project findByProjectId(UUID accountId);
}
