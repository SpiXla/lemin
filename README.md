# **lem-in** - Digital Ant Farm Simulation

## **Project Overview**

The **lem-in** project simulates a digital ant farm where ants move from a starting room (`##start`) to an ending room (`##end`) through a colony of rooms connected by tunnels. The goal is to find the quickest path for the ants to travel from the start to the end while handling various edge cases and input validation.

## **Objectives**

- **Input Handling:** Read data from a file describing rooms, tunnels, and ants.
- **Pathfinding:** Calculate the shortest path or paths for ants to move from `##start` to `##end`.
- **Movement Simulation:** Display the moves made by ants from one room to another in the most efficient way possible.
- **Error Handling:** Handle various invalid or poorly-formatted inputs gracefully.

## **How It Works**

1. **Rooms and Tunnels:**
   - Each room has a name and coordinates (e.g., `Room 1 2`).
   - Rooms are connected by tunnels (e.g., `1-2`).
   - Ants start at `##start` and aim to reach `##end` through the shortest available path.

2. **Ant Movement:**
   - Each ant moves one room at a time.
   - Ants move through tunnels, avoiding traffic jams and ensuring no two ants occupy the same room at the same time (except `##start` and `##end`).

3. **Output:**
   - Display the number of ants, followed by the list of rooms and tunnels.
   - Print the moves of ants in the format `Lx-y`, where `x` is the ant number and `y` is the room the ant is moving to.

## **File Format**

The input file follows this structure:

- **Rooms:** Each room is described by its name and coordinates (e.g., `1 23 3`).
- **Tunnels:** Tunnels are defined by two room numbers connected by a hyphen (e.g., `0-4`).
- **Ants:** The number of ants is given at the top of the file.

Example input file (`test0.txt`):


## **Expected Output Format**

After running the program, the output will show the following:

1. The number of ants.
2. The list of rooms.
3. The list of tunnels.
4. The sequence of movements for the ants.

Example output:


## **How to Run**

1. Clone the repository:

   ```bash
    git clone https://github.com/your-username/lem-in.git
    cd lem-in
    go run . test0.txt
