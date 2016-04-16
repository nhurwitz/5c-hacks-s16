
# coding: utf-8

# In[21]:

import json
data = json.dumps({
  "sideLength": 10,
  "pendingPoints": [
    {
      "x": 7,
      "y": 5,
      "z": 3
    },
    {
      "x": 2,
      "y": 2,
      "z": 2
    }
  ],
  "snakes": {
    "abc123thisIsAnID": {
      "id": "abc123thisIsAnID",
      "color": "#666666",
      "head": {
        "x": 8,
        "y": 8,
        "z": 7
      },
      "tail": [
        {
          "x": 8,
          "y": 7,
          "z": 7
        },
        {
          "x": 8,
          "y": 6,
          "z": 7
        }
      ],
      "direction": "down"
    }
  }
})


# In[67]:

d = json.loads(data)
my_snake_id = 'abc123thisIsAnID'


# In[68]:

state = [0 for i in range(0,d['sideLength']**3)]
state = [[[0 for k in range(0,d['sideLength'])] for j in range(0,d['sideLength'])] for i in range(0,d['sideLength'])]


# In[74]:

for key in d['snakes']:
    for pending_point in d['pendingPoints'][key]['tail']:
            tail_point_x = tail_point['x']
            tail_point_y = tail_point['y']
            tail_point_z = tail_point['z']
            state[tail_point_x][tail_point_y][tail_point_z] = 1
    id =  d['snakes'][key]['id']
    if id == my_snake_id:
        headx = d['snakes'][key]['head']['x']
        heady = d['snakes'][key]['head']['y']
        headz = d['snakes'][key]['head']['z']
        state[headx][heady][headz] = 1
        for tail_point in d['snakes'][key]['tail']:
            tail_point_x = tail_point['x']
            tail_point_y = tail_point['y']
            tail_point_z = tail_point['z']
            state[tail_point_x][tail_point_y][tail_point_z] = 1
    else:
        headx = d['snakes'][key]['head']['x']
        heady = d['snakes'][key]['head']['y']
        headz = d['snakes'][key]['head']['z']
        state[headx][heady][headz] = 2
        for tail_point in d['snakes'][key]['tail']:
            tail_point_x = tail_point['x']
            tail_point_y = tail_point['y']
            tail_point_z = tail_point['z']
            state[tail_point_x][tail_point_y][tail_point_z] = 2


# In[75]:

state


# In[53]:

d


# In[65]:

a = "a"


# In[66]:

b= "b"


# In[ ]:



