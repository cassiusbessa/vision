package com.github.cassiusbessa.vision.domain.service.dtos.auth;

import com.github.cassiusbessa.vision.domain.service.dtos.account.AccountDTO;
import com.github.cassiusbessa.vision.domain.service.dtos.profile.ProfileDTO;

public record MeResponse(
	AccountDTO account,
	ProfileDTO profile,
	String message
) {
	
}
