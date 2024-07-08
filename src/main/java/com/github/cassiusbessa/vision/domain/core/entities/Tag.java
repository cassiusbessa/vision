package com.github.cassiusbessa.vision.domain.core.entities;

import com.github.cassiusbessa.vision.domain.core.valueobjects.TagId;

import java.util.UUID;

public class Tag extends BaseEntity<TagId> {

        private String name;

        public Tag(TagId id, String name) {
            super.setId(id);
            this.name = name;
        }

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }

}
