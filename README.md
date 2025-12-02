# Go hackathon template
## Description
This project is a (backend for the most parts) template to fastly create an applications.

## Service to finish

This template is organized as a set of small services. You can use only what you need for your hackathon project.

0. START OF THE TEMPLATE (Done)

1. **auth-service**  <----ACTIVE
   Handles authentication (signup, login, tokens, password reset).

2. **userprofile-service**  
   Stores user profiles: basic info, preferences, avatars and settings.

3. **crud-service**  
   Generic CRUD endpoints and data-access helpers for your core domain entities.

4. **notification-service**  
   Central layer for dispatching notifications (email, SMS, in-app, etc.).

5. **email-service**  
   Sends transactional emails like verification, password reset and system alerts.

6. **storage-service**  
   Manages file uploads and downloads (images, documents, media).

7. **realtime-service**  
   Provides WebSocket/SSE endpoints for live updates (feeds, dashboards, presence).

8. **search-service**  
   Implements search across users, content and other indexed entities.

9. **dashboard-service**  
   Aggregates metrics and stats for dashboards and overview screens.

10. **chat-service**  
    Enables messaging between users (DMs, group chats, support chat).

11. **payment-service**  
    Integrates with payment providers (e.g. Stripe) to charge users.

12. **billing-service**  
    Deals with plans, subscriptions, invoices and usage-based billing.

13. **calendar-service**  
    Manages events, schedules, reminders and time-based workflows.

14. **access-service**  
    Role- and permission-management (RBAC/ABAC) for protecting resources.

15. **smsotp-service**  
    Sends SMS one-time codes for login, 2FA or phone verification.

16. **analytics-service**  
    Collects and exposes analytics events and basic product metrics.

17. **ml-service**  
    Hosts ML/AI models (inference APIs, feature extraction, scoring).

18. **imageproc-service**  
    Handles image processing (resize, crop, filters, thumbnails).

19. **health-service**  
    Health checks and diagnostics for monitoring and orchestration.


## Packages to finish

1. Core
2. Migrations
3. Infrastructure

## License

This hackathon template is licensed under the **Hackathon Template License (Attribution Required)**.

If you use this template, you **must** credit the author with a note like:

This project uses the Hackathon Template by [Glebegor](https://github.com/Glebegor).
