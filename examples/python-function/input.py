# Schema to validate input against
SCHEMA = {
    'type': 'object',
    'properties': {
        'start_x': {'type': 'number'},
        'start_y': {'type': 'number'},
        'end_x': {'type': 'number'},
        'end_y': {'type': 'number'},
        'world': {
            'type': 'array',
            'items': {
                'type': 'array',
                'items': {'enum': [0, 1]},
                'minItems': 1,
            },
            'minItems': 1,
        },
    },
    'required': [
        'start_x',
        'start_y',
        'end_x',
        'end_y',
        'world',
    ],
}
