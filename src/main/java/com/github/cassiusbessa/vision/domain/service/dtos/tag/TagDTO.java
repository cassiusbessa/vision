package com.github.cassiusbessa.vision.domain.service.dtos.tag;

import java.util.UUID;

public class TagDTO {

        private final UUID id;
        private final String name;


        public TagDTO(UUID id, String name) {
            this.id = id;
            this.name = name;
        }

        public UUID getId() {
            return id;
        }
        public String getName() {
            return name;
        }
}
