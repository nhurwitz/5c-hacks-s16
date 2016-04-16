
# coding: utf-8

# In[25]:

def world_json_to_array(json):
    state = [[[0 for k in range(0,d['sideLength'])] for j in range(0,d['sideLength'])] for i in range(0,d['sideLength'])]
    for key in d['snakes']:
        for pending_point in d['pendingPoints']:
                state[pending_point['x']][pending_point['y']][pending_point['z']] = 3
        id =  d['snakes'][key]['id']
        head = d['snakes'][key]['head']
        if id == my_snake_id:
            state[head['x']][head['y']][head['z']] = 1
            for tail_point in d['snakes'][key]['tail']:
                state[tail_point['x']][tail_point['y']][tail_point['z']] = 1
        else:
            state[head['x']][head['y']][head['z']] = 2
            for tail_point in d['snakes'][key]['tail']:
                 state[tail_point['x']][tail_point['y']][tail_point['z']] = 2
    return [state[i][j][k] for i in range(0,d['sideLength']) for j in range(0,d['sideLength']) for k in range(0,d['sideLength']) ]

