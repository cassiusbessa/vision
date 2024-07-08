package com.github.cassiusbessa.vision.domain.service.mappers;

import com.github.cassiusbessa.vision.domain.core.entities.Account;
import com.github.cassiusbessa.vision.domain.core.entities.Profile;
import com.github.cassiusbessa.vision.domain.core.entities.Tag;
import com.github.cassiusbessa.vision.domain.service.dtos.ProfileCreateCommand;
import org.springframework.stereotype.Component;

import java.util.HashSet;
import java.util.List;

@Component
public class ProfileDataMapper {

    public Profile profileCreateCommandToProfile(ProfileCreateCommand command, Account account, List<Tag> tags) {
        return Profile.builder().
                withName(command.getName()).
                withTitle(command.getTitle()).
                withDescription(command.getDescription()).
                withTechnologies(new HashSet<>(tags)).
                withAccount(account).
                build();
    }
}
