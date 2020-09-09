from pathfinding.core.grid import Grid
from pathfinding.finder.a_star import AStarFinder


def pathfind(world, start_x, start_y, end_x, end_y):
    """Find a path from start to end in a world"""
    grid = Grid(matrix=world)
    start = grid.node(start_x, start_y)
    end = grid.node(end_x, end_y)

    finder = AStarFinder()
    path, runs = finder.find_path(start, end, grid)

    return {
        'path': path,
        'runs': runs,
        'render': grid.grid_str(path=path, start=start, end=end)
    }
