from die import DieCast, RollRequest, get_roll_requests, print_results, roll

def test_die_cast():
    die_cast = DieCast("D6", 4)
    assert die_cast.type == "D6"
    assert die_cast.val == 4

def test_roll_request():
    roll_request = RollRequest("Player 1", 3, 6)
    assert roll_request.name == "Player 1"
    assert roll_request.num == 3
    assert roll_request.sides == 6

def test_get_roll_requests():
    # Assuming some test data
    roll_requests = get_roll_requests()
    assert len(roll_requests) == 2
    assert roll_requests[0].name == "Player 1"
    assert roll_requests[0].num == 2
    assert roll_requests[0].sides == 6
    assert roll_requests[1].name == "Player 2"
    assert roll_requests[1].num == 1
    assert roll_requests[1].sides == 20

def test_print_results(capsys):
    results = [DieCast("D6", 4), DieCast("D6", 2), DieCast("D20", 18)]
    total = 24
    print_results(results, total)
    captured = capsys.readouterr()
    assert captured.out == "Results:\nD6: 4\nD6: 2\nD20: 18\nTotal: 24\n"

def test_roll():
    roll_requests = [
        RollRequest("Player 1", 2, 6),
        RollRequest("Player 2", 1, 20)
    ]
    results = roll(roll_requests)
    assert len(results) == 3
    assert results[0].type == "D6"
    assert results[0].val in range(1, 7)
    assert results[1].type == "D6"
    assert results[1].val in range(1, 7)
    assert results[2].type == "D20"
    assert results[2].val in range(1, 21)
