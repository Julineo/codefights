/*
https://app.codesignal.com/challenge/iWBemn2LDkBokcR6h

An exciting new quantum computing lab has opened up on your university campus. Since there's currently only one quantum computer available, students need to schedule an appointment to use it. Appointments are booked on a first-come first-served basis.

Your friends have found a loophole to book more quantum computing time - every time one of you visits the lab, you book appointments for some of your other friends. To avoid suspicion, you don't all sign everyone up each time; the friends that each one of you signs up is described by the adjacency matrix friends.

friends[i][j] == 1 indicates that the ith friend will book an appointment for the jth friend.
friends[i][j] == 0 means they won't book an appointment for that friend.
When signing up multiple other friends, they're done in ascending order of j values.
Given an integer firstInLine representing the 0th student to have an appointment, your task is to find the kth student to have an appointment.

Notes:

Return -1 if there won't be a kth appointment.
Assume all friends attend all of their booked appointments.
One appointment can be booked per friend, per visit, but there's no limit on how many total appointments any one friend can have booked at a time.
Students can book appointments for themselves.
The adjacency matrix is not symmetric - if friend i books an appointment for friend j, that doesn't necessarily mean friend j would return the favour for friend i.
Example

For

[[1,1,1], 
 [1,0,0], 
 [0,1,1]]
firstInLine = 1, and k = 10, the output should be quantumLabBooking(friends, firstInLine, k) = 2.

Student 1 is first in line, and they sign up student 0.
Student 0 signs up students 0, 1, and 2.
Student 0 signs up students 0, 1, and 2 again.
Student 1 signs up student 0.
Student 2 signs up students 1 and 2.
At this point, 10 appointments have been booked. The order is 1, 0, 0, 1, 2, 0, 1, 2, 0, 1, 2, so the 10th appointment is for student 2.

For

[[0,1,0,0], 
 [0,0,0,0], 
 [1,1,0,0], 
 [1,0,1,0]]
firstInLine = 3, and k = 29, the output should be quantumLabBooking(friends, firstInLine, k) = -1.

Student 3 signs up students 0 and 2.
Student 0 signs up student 1.
Student 2 signs up students 0 and 1.
Student 1 doesn't book any appointments.
Student 0 signs up student 1.
Student 1 doesn't book any appointments.
Student 1 doesn't book any appointments. There are no appointments left.
We make it to the end of the queue without a 29th appointment, so we return -1.
*/

func quantumLabBooking(friends [][]int, firstInLine int, k int) int {
    que := []int{firstInLine}
    cnt := 0
    for cnt < k {
        if cnt >= len(que) { return - 1}
        for i, v := range friends[que[cnt]] {
            if v == 1 {
                que = append(que, i)
            }
        }
        cnt++
    }
    return que[cnt]
}
