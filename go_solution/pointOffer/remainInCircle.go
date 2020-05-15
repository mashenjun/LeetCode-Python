package pointOffer

// 0,1,,n-1这n个数字排成一个圆圈，从数字0开始，每次从这个圆圈里删除第m个数字。求出这个圆圈里剩下的最后一个数字。
// 例如，0、1、2、3、4这5个数字组成一个圆圈，从数字0开始每次删除第3个数字，则删除的前4个数字依次是2、0、4、1，因此最后剩下的数字是3。
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 思路：在数学上，该问题被称为约瑟夫环。有对应的递推公式来计算最后一个留下来的数字的在原始序列中的下标。
// 含义：f(N,M) = 当有N个人，每报到M时杀掉那个人，最终胜利者的在当前序列的编号
// 递推公式：f(N,M) = (f(N−1,M)+M)%N
// 参考：https://blog.csdn.net/u011500062/article/details/72855826
// todo 还未理解
func lastRemaining(n int, m int) int {
	idx := 0
	for i:=2; i<= n; i++ {
		idx = (idx+m) % i
	}
	return idx
}
