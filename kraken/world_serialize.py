
# coding: utf-8

# In[59]:

my_snake_id = 'abc123thisIsAnID'

def world_json_to_array(d):
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

def manhattan_distance(p1, p2):
    p1x = p1['x']
    p1y = p1['y']
    p1z = p1['z']
    p2x = p2['x']
    p2y = p2['y']
    p2z = p2['z']
    return abs(p1x - p2x) + abs(p1y-p2y) + abs(p1z-p2z)

def objective_function(new_state, old_state, id):
    reward = 0;
    isAlive = False
    for key in new_state['snakes']:
        if new_state['snakes'][key]['id'] == my_snake_id:
            isAlive = True
    if not isAlive:
        return 500
    if len(new_state['snakes'][my_snake_id]['tail']) > len(old_state['snakes'][my_snake_id]['tail']):
        reward += 10
    head = new_state['snakes'][my_snake_id]['head']
    minDistance = 500
    for pendingPoint in new_state['pendingPoints']:
        minDistance = min(minDistance, manhattan_distance(head,pendingPoint))
    return reward - minDistance


# In[ ]:




# In[ ]:




# In[ ]:




# In[ ]:




# In[ ]:
