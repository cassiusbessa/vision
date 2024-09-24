package com.github.cassiusbessa.vision.domain.service.ports.output;

import com.github.cassiusbessa.vision.domain.core.entities.Tag;

import java.util.List;
import java.util.UUID;

public interface TagRepository {
    boolean existsAllTags(List<UUID> tagIds);

    List<Tag> findAllById(List<UUID> tagIds);
}
