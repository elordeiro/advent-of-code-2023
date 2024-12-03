card_counts = []

with open("../input.txt") as file:
    card_idx = 0
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
        
        if card_idx == len(card_counts):
            card_counts.append(1)
        
        count = card_counts[card_idx]

        for i in range(card_idx + 1, card_idx+win_count+1):
            if i == len(card_counts):
                card_counts.append(count+1)
            else:
                card_counts[i] += count

        card_idx += 1

print(sum(card_counts))