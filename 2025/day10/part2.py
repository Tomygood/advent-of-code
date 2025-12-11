import gurobipy as gp
from gurobipy import GRB
import numpy as np

tot = 0

input_path = 'C:\\Users/tomlo\\OneDrive\\Bureau\\School\\TELECOM Nancy\\advent-of-code\\2025\\day10\\input.txt'

with open(input_path, 'r') as file:
    for line in file:
        line = line.strip()
        a = line.split(" ")

        buttons = list()

        for button in a[1:-1]:
            buttons.append([int(trig) for trig in button[1:-1].split(",")])

        joltage = [int(j) for j in a[-1][1:-1].split(",")]

        A = list()
        for j in range(len(joltage)):
            A.append([j in button for button in buttons])
        
        A = np.array(A)
        b = np.array(joltage)
        c = np.ones(len(buttons))

        model = gp.Model()
        model.Params.LogToConsole = 0

        x = model.addMVar(shape=c.shape, vtype=GRB.INTEGER, name="x")

        model.setObjective(c @ x, GRB.MINIMIZE)

        model.addConstr(A @ x == b)

        model.optimize()
        tot += model.ObjVal


print("Part 2 result:", int(tot))