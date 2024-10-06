# memory
A set of services that support the Memory App.

Another name of the application will be Twinkle.

## Conception
As a father, I want to cherish the moments that I have with my child. In addition, I love creating new solutions to problems. One of the problems that I want to solve is to have a way to remember every special moment with my child. Thinking about that, I thought that it would be a great idea to create an application that will send me a text message daily. The text message will ask me for a significant memory of the day with my child. I can send a response back with a picture or text. The response will be stored so I can recall it in the future.

Storing the memory will allow me to do the following:
#. Recall significant and random events with my child.
#. Give me the opportunity to create a time capsule for my child.

## Design
The application will consist of several parts.
#. There will be an API that will be the interface for the database and blob store
#. There will be an event engine that will manage asynchronous events and cron events.
