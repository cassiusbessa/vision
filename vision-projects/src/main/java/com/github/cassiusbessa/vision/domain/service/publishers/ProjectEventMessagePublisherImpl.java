package com.github.cassiusbessa.vision.domain.service.publishers;

import com.github.cassiusbessa.vision.domain.core.events.ProjectCreatedEvent;
import com.github.cassiusbessa.vision.domain.core.events.ProjectDeletedEvent;
import com.github.cassiusbessa.vision.domain.core.events.ProjectEvent;
import com.github.cassiusbessa.vision.domain.core.events.ProjectUpdatedEvent;
import com.github.cassiusbessa.vision.domain.service.ports.output.messages.ProjectEventMessagePublisher;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;
import org.json.JSONObject;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.URI;
import java.net.URL;
import java.nio.charset.StandardCharsets;

@Slf4j
@Component
public class ProjectEventMessagePublisherImpl implements ProjectEventMessagePublisher {

    @Override
    public void publish(ProjectCreatedEvent event) {
        log.info("Project created event fired: {}", event.getProject().getTitle());
        sendPostRequest(event);
    }

    @Override
    public void publish(ProjectUpdatedEvent event) {
        log.info("Project updated event fired: {}", event.getProject().getTitle());
        sendPutRequest(event);
    }

    @Override
    public void publish(ProjectDeletedEvent event) {
        log.info("Project deleted event fired: {}", event.getProject().getTitle());
        sendDeleteRequest(event);
    }

    @Override
    public void publish(ProjectEvent event) {
        log.info("Project event: {}", event.getProject().getTitle());
        sendPostRequest(event);
    }

    private void sendPostRequest(ProjectEvent event) {
        String url = "http://localhost:8888/posts";
        HttpURLConnection connection = null;
        try {
            JSONObject json = new JSONObject();
            json.put("project_id", event.getProject().getId().getValue().toString());
            json.put("author_id", event.getProject().getAccount().getId().getValue().toString());
            json.put("title", event.getProject().getTitle());
            json.put("content", event.getProject().getDescription());

            URL endpoint = new URI(url).toURL();
            connection = (HttpURLConnection) endpoint.openConnection();
            connection.setRequestMethod("POST");
            connection.setRequestProperty("Content-Type", "application/json; utf-8");
            connection.setDoOutput(true);

            try (OutputStream os = connection.getOutputStream()) {
                byte[] input = json.toString().getBytes(StandardCharsets.UTF_8);
                os.write(input, 0, input.length);
            }

            int responseCode = connection.getResponseCode();
            log.info("POST Post Response Code: {}", responseCode);

            if (responseCode == HttpURLConnection.HTTP_CREATED) {
                log.info("POST was successful.");
            } else {
                log.warn("POST request did not work. Response Code: {}", responseCode);
            }

        } catch (Exception e) {
            log.error("Error sending POST request", e);
        } finally {
            if (connection != null) {
                connection.disconnect();
            }
        }
    }

    private void sendPutRequest(ProjectEvent event) {
        String url = "http://localhost:8888/posts";
        HttpURLConnection connection = null;
        try {
            JSONObject json = new JSONObject();
            json.put("project_id", event.getProject().getId().getValue().toString());
            json.put("title", event.getProject().getTitle());
            json.put("content", event.getProject().getDescription());
            json.put("repo_link", event.getProject().getLinks().getRepository());
            json.put("demo_link", event.getProject().getLinks().getDemo());

            URL endpoint = new URI(url).toURL();
            connection = (HttpURLConnection) endpoint.openConnection();
            connection.setRequestMethod("PUT");
            connection.setRequestProperty("Content-Type", "application/json; utf-8");
            connection.setDoOutput(true);

            try (OutputStream os = connection.getOutputStream()) {
                byte[] input = json.toString().getBytes(StandardCharsets.UTF_8);
                os.write(input, 0, input.length);
            }

            int responseCode = connection.getResponseCode();
            log.info("Put Post Response Code: {}", responseCode);

            if (responseCode == HttpURLConnection.HTTP_OK) {
                log.info("PUT was successful.");
            } else {
                log.warn("PUT request did not work. Response Code: {}", responseCode);
            }

        } catch (Exception e) {
            log.error("Error sending PUT request", e);
        } finally {
            if (connection != null) {
                connection.disconnect();
            }
        }
    }

    private void sendDeleteRequest(ProjectEvent event) {
        String url = "http://localhost:8888/posts/";
        HttpURLConnection connection = null;
        try {
            JSONObject json = new JSONObject();
            json.put("project_id", event.getProject().getId().getValue().toString());
            json.put("author_id", event.getProject().getAccount().getId().getValue().toString());
            json.put("title", event.getProject().getTitle());
            json.put("content", event.getProject().getDescription());

            URL endpoint = new URI(url).toURL();
            connection = (HttpURLConnection) endpoint.openConnection();
            connection.setRequestMethod("PUT");
            connection.setRequestProperty("Content-Type", "application/json; utf-8");
            connection.setDoOutput(true);

            try (OutputStream os = connection.getOutputStream()) {
                byte[] input = json.toString().getBytes(StandardCharsets.UTF_8);
                os.write(input, 0, input.length);
            }

            int responseCode = connection.getResponseCode();
            log.info("DELETE Post Response Code: {}", responseCode);

            if (responseCode == HttpURLConnection.HTTP_OK) {
                log.info("DELETE was successful.");
            } else {
                log.warn("DELETE request did not work. Response Code: {}", responseCode);
            }

        } catch (Exception e) {
            log.error("Error sending PUT request", e);
        } finally {
            if (connection != null) {
                connection.disconnect();
            }
        }
    }


}
