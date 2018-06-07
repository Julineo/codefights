/*
Mr. Hudson is teaching a course on algorithms at computer school. He's assigning his class a project that involves a group presentation. Each presentation will happen one at a time and the groups can have as many as groupSize members.

Based on prior experience, Mr. Hudson has a list of estimates for each of the students, representing the minimum number of minutes that a presentation will take if that student is part of the group. These time estimates are stored in an array timeEstimates (For confidentiality reasons, the students' names have been anonymized, so it's just the number of minutes each will take).

Mr. Hudson wants to take up as little class time as possible for the whole class to give their presentations. Return the total minimum number of minutes the presentations will be expected to take if the groups are assigned in such a way that the estimated total time is minimized.

Example

For timeEstimates = [12, 7, 3, 6, 5, 6, 10, 3, 9, 6] and groupSize = 3, the output should be assignGroups(timeEstimates, groupSize) = 28.
*/

import "sort"

func assignGroups(timeEstimates []int, groupSize int) (r int) {
    sort.Ints(timeEstimates)
    for i := len(timeEstimates) - 1; i >= 0; i = i - groupSize  {
        r = r + timeEstimates[i]
    }
    return r
}
