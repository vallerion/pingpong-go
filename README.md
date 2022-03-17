# pingpong-go
PIngPong game based on ebiten, written in go.
![Game Process](/client/assets/game_process.gif)

## Todo

- [ ] add background picture
- [ ] handle exit
- [ ] server for multiplayer

## Multiplayer flow
1. One player init game (left player). 
2. Create connection with server
3. Check if room exists connect to it, otherwise create room
4. Waiting for another player
5. Start the game, room change status from PENDING to PLAYING
6. Positions of leftPlayer, rightPlayer and ball serve by server