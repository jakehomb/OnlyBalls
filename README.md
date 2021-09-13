# OnlyBalls

### API Functionality

- [ ] Login function
- [ ] Database connectivity
- [ ] Ball upload functionality
    - [ ] Validate the size of the balls per the DBMS requirements
    - [ ] If needed, compress the balls
    - [ ] Hash of Balls or UUID 
    - [ ] User ID associated
    - [ ] base64 encoding of balls for storage in DBMS
- [ ] Ball retrieval route 
    - [ ] Balls hash should be included in route URI
    - [ ] Balls hash used to query DBMS for Ball data
    - [ ] Route Base64 decodes image file
    - [ ] Return byte array to caller as the image
- [ ] 