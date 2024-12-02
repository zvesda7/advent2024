using System.Runtime.Serialization;

string[] lines = File.ReadAllLines(@"..\..\..\input.txt");
var parsed = lines.Select(x => x.Split(' ').Select(x => int.Parse(x)));


var safe1 = parsed.Count(x => safe(x));
Console.WriteLine($"Part1 {safe1}");

var safe2 = parsed.Count(line => {        
    var safes = line.Select((excVals, excI) => safe(line.Where((x,i) => i != excI))).Count(x => x);
    return safes > 0 || safe(line);
});
Console.WriteLine($"Part2 {safe2}");

public static partial class Program
{    
    public static bool safe(IEnumerable<int> x)
    {
        var deltas = x.Zip(x.Skip(1), (x, y) => y - x);    
        var ascGood = deltas.Count(x => x > 0 && Math.Abs(x) <= 3);
        var descGood = deltas.Count(x => x < 0 && Math.Abs(x) <= 3);    
        return ascGood == deltas.Count() || descGood == deltas.Count();
    }
}