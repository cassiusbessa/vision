package com.github.cassiusbessa.vision.dataaccess.mappers;

import com.github.cassiusbessa.vision.dataaccess.entities.TagDataBaseEntity;
import com.github.cassiusbessa.vision.domain.core.entities.Tag;
import com.github.cassiusbessa.vision.domain.core.valueobjects.TagId;
import org.springframework.stereotype.Component;

@Component
public class TagDataBaseMapper {

    public Tag dbEntityToTag(TagDataBaseEntity dbEntity) {
        return new Tag(new TagId(dbEntity.getId()), dbEntity.getName());
    }

    public TagDataBaseEntity tagToDbEntity(Tag tag) {
        return new TagDataBaseEntity(tag.getId().getValue(), tag.getName());
    }
}
