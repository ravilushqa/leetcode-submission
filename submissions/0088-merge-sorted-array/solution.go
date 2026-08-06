func merge(nums1 []int, m int, nums2 []int, n int)  {
    writeIdx := len(nums1) - 1
    read1Idx := m - 1
    read2Idx := n - 1

    // >=!
    for read2Idx >= 0 {
        // >=!
        if read1Idx >= 0 && nums1[read1Idx] > nums2[read2Idx] {
            nums1[writeIdx] = nums1[read1Idx]
            read1Idx--
        } else {
            nums1[writeIdx] = nums2[read2Idx]
            read2Idx--
        }

        writeIdx--
    }
}
