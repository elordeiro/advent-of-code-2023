sum = 0

with open("../input.txt") as file:
    while (line := file.readline()):
        parts = line.split("|")
        winners = sorted(map(int, parts[0][parts[0].index(":")+1:].strip().split()))
        nums = sorted(map(int, parts[1].strip().split()))

        win_count = 0
        i, j = 0, 0
        while i < len(winners) and j < len(nums):
            if winners[i] == nums[j]:
                win_count += 1
                i += 1
                j += 1
                continue
            if winners[i] > nums[j]:
                j += 1
            else: 
                i += 1
        
        sum += int(2 ** (win_count-1))

print(sum)