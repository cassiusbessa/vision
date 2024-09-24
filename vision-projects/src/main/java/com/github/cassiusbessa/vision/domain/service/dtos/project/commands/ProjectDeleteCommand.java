package com.github.cassiusbessa.vision.domain.service.dtos.project.commands;

import java.util.UUID;

public record ProjectDeleteCommand(UUID projectId, UUID accountId) {
}
