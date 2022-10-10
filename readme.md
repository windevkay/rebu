**REBU**

_About this project_

REBU is a microservice that simulates the process of matching a driver to a rider that has made a request (on UBER), in as optimal a way as possible.

_Objectives_

- Compute rider radius (in an incremental fashion), and represent rider/driver(s) as nodes snapshot. e.g start at a 1 mile radius and go from there depending on driver availability within the radius (in reality this data can probably come from google maps).
- Apply an optimal algorithm to compute:
  - Shortest path between driver/rider
  - Filter occupied vs unoccupied drivers
  - Determine if an occupied driver in the process of completing a trip mignt be more optimal than an unoccupied driver that is much farther away.
- Return selected/matched driver details
