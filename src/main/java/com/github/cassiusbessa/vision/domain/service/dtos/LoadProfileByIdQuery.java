package com.github.cassiusbessa.vision.domain.service.dtos;

import java.util.UUID;

public class LoadProfileByIdQuery {

    private final UUID id;

    public LoadProfileByIdQuery(UUID id) {
        this.id = id;
    }

    public UUID getId() {
        return id;
    }
}
